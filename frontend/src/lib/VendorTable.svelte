<script lang="ts">
  import type { Vendor } from "../types.ts";

  let { vendor }: { vendor: Vendor } = $props();
</script>

<div>
  <h1>
    {vendor.name}
  </h1>
  <table>
    <thead>
      <tr>
        <th>Item Name</th>
        <th>Type</th>
        <th>Standing</th>
        <th>Volume</th>
        <th>Average Price</th>
      </tr>
    </thead>
    <tbody>
      {#each vendor.items as item}
        <tr>
          <td>{item.name}</td>
          <td>{item.type}</td>
          <td>{item.standing}</td>
          <td>{item.volume}</td>
          <td>{item.weightedPrice}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
