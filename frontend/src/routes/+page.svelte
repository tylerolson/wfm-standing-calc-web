<script lang="ts">
  import { type BasicVendor, type BasicVendorsResponse } from "$lib/types";
  import BasicVendorPanel from "$lib/BasicVendorPanel.svelte";
  import { onMount } from "svelte";

  let loadStatus = $state("Loading...");
  let vendors = $state<BasicVendor[]>([]);
  let updatedAt = $state<Date>();
  let updating = $state(false);

  onMount(async () => {
    const response = await fetch("/api/vendors");

    if (!response.ok) {
      loadStatus = `HTTP error ${response.status} (${response.statusText}) The backend server may be down, try again later.`;
      return;
    }

    const data: BasicVendorsResponse = await response.json();

    vendors = data.vendors;
    updatedAt = new Date(data.updatedAt);
    updating = data.updating;
  });
</script>

<svelte:head>
  <title>WFM Calculator</title>
</svelte:head>

{#if vendors.length === 0}
  <div class="m-6">
    <p class="text-3xl text-gray-100">{loadStatus}</p>
  </div>
{:else}
  <div class="mt-1">
    <div class="flex justify-center">
      {#if updating}
        <p class="pb-3 text-gray-100">Updating...</p>
      {:else}
        <p class="pb-3 text-gray-100">Updated at: {updatedAt?.toLocaleString()}</p>
      {/if}
    </div>

    <div
      class="mx-auto grid w-full max-w-[1500px] auto-cols-fr grid-cols-1 gap-8 p-4 sm:grid-cols-2 xl:grid-cols-3"
    >
      {#each vendors as vendor (vendor.slug)}
        <a href={`/vendors/${vendor.slug}`}>
          <BasicVendorPanel {vendor}></BasicVendorPanel>
        </a>
      {/each}
    </div>
  </div>
{/if}
