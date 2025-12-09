<script lang="ts">
  import { onMount } from "svelte";
  import { page } from "$app/state";
  import VendorTable from "$lib/VendorTable.svelte";
  import type { Vendor, VendorResponse } from "$lib/types";

  let loadStatus = $state("Loading...");
  let vendor = $state<Vendor>();
  let updatedAt = $state<Date>();
  let updating = $state(false);

  let filterText = $state("");
  onMount(async () => {
    const response = await fetch(`/api/vendors/${page.params.slug}`);

    if (!response.ok) {
      if (response.status == 404) {
        loadStatus = `Vendor "${page.params.slug}" not found`;
      } else {
        loadStatus = `HTTP error ${response.status} (${response.statusText}) The backend server may be down, try again later.`;
      }
      return;
    }

    const data: VendorResponse = await response.json();

    console.log(data);
    vendor = data.vendor;
    updatedAt = new Date(data.updatedAt);
    updating = data.updating;
  });
</script>

{#if vendor === undefined}
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

    <div class="mx-auto grid w-full max-w-[1500px] auto-cols-fr grid-cols-1 gap-8 px-4 py-4">
      <!-- svelte-ignore a11y_autofocus -->
      <input
        placeholder="Search"
        class=" rounded-md bg-gray-700 px-6 py-3 text-gray-100 shadow-xl focus:outline-gray-400"
        autofocus
        bind:value={filterText}
      />
      <VendorTable {vendor} {filterText}></VendorTable>
    </div>
  </div>
{/if}
