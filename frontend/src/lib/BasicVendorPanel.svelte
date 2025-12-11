<script lang="ts">
  import type { BasicVendor } from "$lib/types";
  import VendorIcon from "$lib/VendorIcon.svelte";
  let { vendor }: { vendor: BasicVendor } = $props();
</script>

<div class="h-full rounded-lg bg-[#1F2937] text-[#F9FAFB] shadow-lg">
  <div class={`flex items-center rounded-t-lg py-4 pl-4 text-3xl ${vendor.slug}`}>
    <VendorIcon {vendor}></VendorIcon>
    {vendor.name}
  </div>
  <div class="p-6">
    <a href={`https://warframe.market/items/${vendor.mostProfitable.name}?type=sell`}>
      <p class="text-xl">
        Best Profit: {vendor.mostProfitable.name}
      </p>
      <p class="mb-2 text-xl text-[#9CA3AF]">
        {vendor.mostProfitable.price.toFixed(2)} platinum {vendor.mostProfitable.volume.toFixed(2)} volume
      </p>
    </a>

    <a href={`https://warframe.market/items/${vendor.mostVolume.name}?type=sell`}>
      <p class="text-xl">Most Volume: {vendor.mostVolume.name}</p>
      <p class="mb-2 text-xl text-[#9CA3AF]">
        {vendor.mostVolume.price.toFixed(2)} platinum {vendor.mostVolume.volume.toFixed(2)} volume
      </p>
    </a>

    <a href={`https://warframe.market/items/${vendor.mostEfficient.name}?type=sell`}>
      <p class="text-xl">Most Efficient: {vendor.mostEfficient.name}</p>
      <p class="mb-2 text-xl text-[#9CA3AF]">
        {vendor.mostEfficient.price.toFixed(2)} platinum {vendor.mostEfficient.volume.toFixed(2)} volume
      </p>
    </a>
  </div>
</div>

<style>
</style>
