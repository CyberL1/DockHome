<script lang="ts">
  import type { Props } from "./types";
  import { tick } from "svelte";

  let { id = "menu", data, items }: Props = $props();

  let menuVisible = $state(false);
  let menuX = $state(0);
  let menuY = $state(0);

  // svelte-ignore non_reactive_update
  let menu: HTMLElement;

  export async function open(event: MouseEvent, menuData: object) {
    const rect = (event.target as HTMLElement).getBoundingClientRect();

    data = menuData;
    menuVisible = true;

    await tick();

    const menuRect = menu.getBoundingClientRect();

    let x = event.clientX;
    let y = rect.bottom + window.scrollY;

    if (x + menuRect.width > window.innerWidth) {
      x = window.innerWidth - menuRect.width - 10;
    }

    if (y + menuRect.height > window.innerHeight + window.scrollY) {
      y = rect.top + window.scrollY - menuRect.height;
    }

    menuX = x;
    menuY = y;
  }

  export function close() {
    menuVisible = false;
  }
</script>

{#if menuVisible}
  <div bind:this={menu} {id} style="left: {menuX}px; top: {menuY}px;">
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
