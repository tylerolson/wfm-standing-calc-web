import type { components } from "$lib/api";
import type { PageLoad } from "./$types";

type BasicVendorsResponse = components["schemas"]["BasicVendorsResponseBody"];

export const load: PageLoad = async ({ fetch }) => {
  const response = await fetch("/api/vendors");

  if (!response.ok) {
    return {
      vendors: [],
      updatedAt: null,
      updating: false,
      error: `HTTP error ${response.status} (${response.statusText}) The backend server may be down, try again later.`,
    };
  }

  const data: BasicVendorsResponse = await response.json();

  return {
    vendors: data.vendors,
    updatedAt: new Date(data.updatedAt),
    updating: data.updating,
    error: null,
  };
};
