<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/state";
  import type { Vendor } from "$lib/types";
  import { ItemType } from "$lib/types";
  import VendorIcon from "$lib/VendorIcon.svelte";

  let { vendor, filterText }: { vendor: Vendor; filterText: string } = $props();
  type SortKey = "name" | "type" | "standing" | "volume" | "price";
  type SortDirection = "asc" | "desc" | null;

  let sortKey = $state<SortKey | null>((page.url.searchParams.get("sortKey") as SortKey) || null);
  let sortDirection = $state<SortDirection>(
    (page.url.searchParams.get("sortDirection") as SortDirection) || null,
  );

  function updateURL() {
    const params = new URLSearchParams(page.url.searchParams);

    if (sortKey && sortDirection) {
      params.set("sortKey", sortKey);
      params.set("sortDirection", sortDirection);
    } else {
      params.delete("sortKey");
      params.delete("sortDirection");
    }

    goto(`?${params.toString()}`, { keepFocus: true, noScroll: true, replaceState: true });
  }

  function handleSort(key: SortKey) {
    if (sortKey === key) {
      // Toggle through: asc -> desc -> null
      if (sortDirection === "desc") {
        sortDirection = "asc";
      } else if (sortDirection === "asc") {
        sortDirection = null;
        sortKey = null;
      }
    } else {
      sortKey = key;
      sortDirection = "desc";
    }

    updateURL();
  }

  let sortedItems = $derived(() => {
    let filtered = vendor.items.filter(
      (item) => item.name.toLowerCase().includes(filterText.toLowerCase()) || filterText === "",
    );

    if (!sortKey || !sortDirection) return filtered;

    return [...filtered].sort((a, b) => {
      let aVal: string | number = "";
      let bVal: string | number = "";

      switch (sortKey) {
        case "name":
          aVal = a.name.toLowerCase();
          bVal = b.name.toLowerCase();
          break;
        case "type":
          aVal = ItemType[a.type];
          bVal = ItemType[b.type];
          break;
        case "standing":
          aVal = a.standing;
          bVal = b.standing;
          break;
        case "volume":
          aVal = a.volume;
          bVal = b.volume;
          break;
        case "price":
          aVal = a.price;
          bVal = b.price;
          break;
      }

      if (aVal < bVal) return sortDirection === "asc" ? -1 : 1;
      if (aVal > bVal) return sortDirection === "asc" ? 1 : -1;
      return 0;
    });
  });

  function getSortIcon(key: SortKey) {
    if (sortKey !== key) return "↕";
    return sortDirection === "asc" ? "↑" : "↓";
  }

  // let hidden = $state(false);
  // let tBody: HTMLTableSectionElement;
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
        <th
          scope="col"
          class="cursor-pointer py-3 pl-4 select-none hover:text-gray-200"
          onclick={() => handleSort("name")}
        >
          Item Name <span class="ml-1">{getSortIcon("name")}</span>
        </th>
        <th
          scope="col"
          class="cursor-pointer py-3 pl-4 select-none hover:text-gray-200"
          onclick={() => handleSort("type")}
        >
          Type <span class="ml-1">{getSortIcon("type")}</span>
        </th>
        <th
          scope="col"
          class="cursor-pointer py-3 pl-4 select-none hover:text-gray-200"
          onclick={() => handleSort("standing")}
        >
          Standing <span class="ml-1">{getSortIcon("standing")}</span>
        </th>
        <th
          scope="col"
          class="cursor-pointer py-3 pl-4 select-none hover:text-gray-200"
          onclick={() => handleSort("volume")}
        >
          Volume <span class="ml-1">{getSortIcon("volume")}</span>
        </th>
        <th
          scope="col"
          class="cursor-pointer px-4 py-3 select-none hover:text-gray-200"
          onclick={() => handleSort("price")}
        >
          Avg Price <span class="ml-1">{getSortIcon("price")}</span>
        </th>
      </tr>
    </thead>
    <tbody>
      {#each sortedItems() as item}
        <tr class="border-b border-gray-700 bg-gray-800 transition-all hover:bg-gray-700">
          <td class="py-4 pl-4 font-medium whitespace-nowrap text-gray-100">
            <a
              href={`https://warframe.market/items/${item.name}?type=sell`}
              target="_blank"
              rel="noopener noreferrer"
            >
              {item.name}
            </a>
          </td>
          <td class="py-4 pl-4">{ItemType[item.type]}</td>
          <td class="py-4 pl-4">{item.standing}</td>
          <td class="py-4 pl-4">{item.volume.toFixed(2)}</td>
          <td class="px-4 py-4">{item.price.toFixed(2)}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
