<script lang="ts">
  import type { Props } from "./types";

  let { id = "menu", data, items }: Props = $props();

  let menuVisible = $state(false);
  let menuX = $state(0);
  let menuY = $state(0);

  export function open(event: MouseEvent, menuData: object) {
    const rect = (event.target as HTMLElement).getBoundingClientRect();
    menuX = event.clientX;
    menuY = rect.bottom + window.scrollY;

    data = menuData;
    menuVisible = true;
  }

  export function close() {
    menuVisible = false;
  }
</script>

{#if menuVisible}
  <div {id} style="left: {menuX}px; top: {menuY}px;">
    {#each items as item}
      <button onclick={item.onclick}>{item.title}</button>
    {/each}
  </div>
{/if}

<style>
  div {
    position: absolute;
    background: white;
    border: 1px solid #ccc;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    min-width: 150px;
    border-radius: 8px;
    user-select: none;
    z-index: 1000;
  }

  div button {
    cursor: pointer;
    display: flex;
    background: none;
    border: none;
    width: 100%;
    padding: 10px 16px;
    font-size: 1em;
  }

  div button:hover {
    background-color: lightgrey;
  }
</style>
