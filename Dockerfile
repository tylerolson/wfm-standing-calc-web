# Frontend
FROM node:22.15.0-alpine3.21 AS frontend-builder
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

# Cache dependencies
COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

# Copy rest of frontend
COPY frontend ./
RUN pnpm run build

# Backend
FROM golang:1.24-alpine3.21 AS backend-builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
COPY --from=frontend-builder /app/dist ./frontend/dist

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server


FROM gcr.io/distroless/static-debian12:latest
WORKDIR /app
COPY --from=backend-builder /app/server ./
CMD ["./server"]
