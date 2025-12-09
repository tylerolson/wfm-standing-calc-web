<script lang="ts">
  import type { Vendor } from "$lib/types";
  import { ItemType } from "$lib/types";
  import VendorIcon from "$lib/VendorIcon.svelte";

  let hidden = $state(false);
  let tBody: HTMLTableSectionElement;

  let { vendor, filterText }: { vendor: Vendor; filterText: string } = $props();

  // $effect(() => {
  //   // Call this because we need to call the effect when that updates
  //   // We don't need to use it because the logic is handled in the HTML
  //   filterText;
  //   hidden = tBody.rows.length == 0;
  // });
</script>

<!-- I tried to use #if here, but when the object gets deleted I can't do the filter logic -->
<div class={`mt-4 h-full overflow-x-auto  rounded-lg bg-[#1F2937] text-[#F9FAFB] shadow-md`}>
  <table class="w-full text-left text-sm text-gray-400 rtl:text-right">
    <caption class={`py-4 pl-4 text-left text-3xl ${vendor.slug}`}>
      <div class="flex items-center">
        <VendorIcon {vendor}></VendorIcon>
        {vendor.name}
      </div>
    </caption>
    <thead class="bg-gray-700 text-xs text-gray-400 uppercase">
      <tr>
        <th scope="col" class="py-3 pl-4">Item Name</th>
        <th scope="col" class="py-3 pl-4">Type</th>
        <th scope="col" class="py-3 pl-4">Standing</th>
        <th scope="col" class="py-3 pl-4">Volume</th>
        <th scope="col" class="px-4 py-3">Avg Price</th>
      </tr>
    </thead>
    <tbody bind:this={tBody}>
      {#each vendor.items as item}
        {#if item.name.toLowerCase().includes(filterText.toLowerCase()) || filterText === ""}
          <tr class="border-b border-gray-700 bg-gray-800 transition-all hover:bg-gray-700">
            <td class="py-4 pl-4 font-medium whitespace-nowrap text-gray-100">{item.name}</td>
            <td class="py-4 pl-4">{ItemType[item.type]}</td>
            <td class="py-4 pl-4">{item.standing}</td>
            <td class="py-4 pl-4">{item.volume.toFixed(2)}</td>
            <td class="px-4 py-4">{item.price.toFixed(2)}</td>
          </tr>
        {/if}
      {/each}
    </tbody>
  </table>
</div>
