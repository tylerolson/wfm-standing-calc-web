<script lang="ts">
  import type { BasicVendor } from "$lib/types";
  import arbitierIcon from "$lib/assets/ArbitarIconGrey.png";
  import cephalonIcon from "$lib/assets/CephalonIconLightGold.png";
  import newLokaIcon from "$lib/assets/LokaIconGreen.png";
  import redVeilIcon from "$lib/assets/RedVeilIconLightRed.png";
  import steelIcon from "$lib/assets/SteelIconPink.png";
  import perrinIcon from "$lib/assets/PerrinSequenceIconBlue.png";

  let { vendor }: { vendor: BasicVendor } = $props();

  const ICONS = {
    arbiters_of_hexis: arbitierIcon,
    cephalon_suda: cephalonIcon,
    new_loka: newLokaIcon,
    red_veil: redVeilIcon,
    steel_meridian: steelIcon,
    the_perrin_sequence: perrinIcon,
  };

  console.log(vendor);
</script>

<div class="h-full rounded-lg bg-[#1F2937] text-[#F9FAFB] shadow-md">
  <div class={`flex items-center rounded-t-lg py-4 pl-4 text-3xl antialiased ${vendor.slug}`}>
    <img src={ICONS[vendor.slug]} alt="Steel meridian icon" class="mr-4 w-16" />
    {vendor.name}
  </div>
  <div class="p-6">
    <p class="text-xl">Best Profit: {vendor.mostProfitable.name}</p>
    <p class="mb-2 text-xl text-[#9CA3AF]">
      {vendor.mostProfitable.price.toFixed(2)} platinum {vendor.mostProfitable.volume.toFixed(2)} volume
    </p>

    <p class="text-xl">Most Volume: {vendor.mostVolume.name}</p>
    <p class="mb-2 text-xl text-[#9CA3AF]">
      {vendor.mostVolume.price.toFixed(2)} platinum {vendor.mostVolume.volume.toFixed(2)} volume
    </p>

    <p class="text-xl">Most Efficient: {vendor.mostEfficient.name}</p>
    <p class="mb-2 text-xl text-[#9CA3AF]">
      {vendor.mostEfficient.price.toFixed(2)} platinum {vendor.mostEfficient.volume.toFixed(2)} volume
    </p>
  </div>
</div>

<style>
  .arbiters_of_hexis {
    @apply border-b-2 border-[#CFE1E3] bg-[#2F3335] text-[#CFE1E3];
  }
  .cephalon_suda {
    @apply border-b-2 border-[#F8FDAC] bg-[#312C4A] text-[#F8FDAC];
  }
  .new_loka {
    @apply border-b-2 border-[#C1FFBD] bg-[#223025] text-[#C1FFBD];
  }
  .red_veil {
    @apply border-b-2 border-[#FD8A87] bg-[#31132E] text-[#FD8A87];
  }
  .steel_meridian {
    @apply border-b-2 border-[#F8B283] bg-[#233238] text-[#F8B283];
  }

  .the_perrin_sequence {
    @apply border-b-2 border-[#91DBFF] bg-[#313A4F] text-[#91DBFF];
  }
</style>
