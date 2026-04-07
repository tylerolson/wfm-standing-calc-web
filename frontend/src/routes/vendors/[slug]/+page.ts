import type { components } from "$lib/api";
import type { PageLoad } from "./$types";

type VendorResponse = components["schemas"]["VendorsResponseBody"];

export const load: PageLoad = async ({ fetch, params }) => {
  const response = await fetch(`/api/vendors/${params.slug}`);

  if (!response.ok) {
    if (response.status == 404) {
      return {
        error: `Vendor "${params.slug}" not found`,
      };
    } else {
      return {
        error: `HTTP error ${response.status} (${response.statusText}) The backend server may be down, try again later.`,
      };
    }
  }

  const data: VendorResponse = await response.json();

  return {
    vendor: data.vendor,
    updatedAt: new Date(data.updatedAtMs),
    updating: data.updating,
    error: null,
  };
};
