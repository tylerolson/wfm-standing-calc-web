<script lang="ts">
  import type { BasicVendor, Vendor, VendorSlug } from "$lib/types";

  let { vendor }: { vendor: Vendor | BasicVendor } = $props();

  import arbitierIcon from "$lib/assets/ArbitarIconGrey.png";
  import cephalonIcon from "$lib/assets/CephalonIconLightGold.png";
  import newLokaIcon from "$lib/assets/LokaIconGreen.png";
  import redVeilIcon from "$lib/assets/RedVeilIconLightRed.png";
  import steelIcon from "$lib/assets/SteelIconPink.png";
  import perrinIcon from "$lib/assets/PerrinSequenceIconBlue.png";

  export const ICONS: Record<VendorSlug, string> = {
    arbiters_of_hexis: arbitierIcon,
    cephalon_suda: cephalonIcon,
    new_loka: newLokaIcon,
    red_veil: redVeilIcon,
    steel_meridian: steelIcon,
    the_perrin_sequence: perrinIcon,
  };
</script>

<img src={ICONS[vendor.slug]} alt={`${vendor.name} icon`} class="mr-4 w-16" />
