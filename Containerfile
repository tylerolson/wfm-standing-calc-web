# Frontend
FROM node:22.21-alpine AS frontend-builder
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

# Cache dependencies
COPY frontend/pnpm-lock.yaml frontend/package.json frontend/pnpm-workspace.yaml* frontend/.npmrc* ./
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
COPY --from=frontend-builder /app/build ./frontend/build

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server


FROM gcr.io/distroless/static-debian12:latest
WORKDIR /app
COPY --from=backend-builder /app/server ./
CMD ["./server"]
