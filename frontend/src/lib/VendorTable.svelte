<script lang="ts">
  import type { Vendor } from "../types";
  import { ItemType } from "../types";

  let { vendor }: { vendor: Vendor } = $props();
</script>

<div class="overflow-x-auto shadow-md rounded-xl bg-gray-800 mt-4">
  <table class="w-full text-sm text-left rtl:text-right text-gray-400">
    <caption
      class="pl-4 pt-4 pb-2 text-lg font-semibold text-left rtl:text-right text-white bg-gray-700"
    >
      {vendor.name}
    </caption>
    <thead class="text-xs uppercase bg-gray-700 text-gray-400">
      <tr>
        <th scope="col" class="pl-4 py-3">Item Name</th>
        <th scope="col" class="pl-4 py-3">Type</th>
        <th scope="col" class="pl-4 py-3">Standing</th>
        <th scope="col" class="pl-4 py-3">Volume</th>
        <th scope="col" class="px-4 py-3">Avg Price</th>
      </tr>
    </thead>
    <tbody>
      {#each vendor.items as item}
        <tr class=" border-b bg-gray-800 border-gray-700">
          <td class="pl-4 py-4 font-medium text-gray-100 whitespace-nowrap">{item.name}</td>
          <td class="pl-4 py-4">{ItemType[item.type]}</td>
          <td class="pl-4 py-4">{item.standing}</td>
          <td class="pl-4 py-4">{item.volume.toFixed(2)}</td>
          <td class="px-4 py-4">{item.weightedPrice.toFixed(2)}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
