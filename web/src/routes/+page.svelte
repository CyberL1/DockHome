<script lang="ts">
  import { page } from "$app/state";
  import type { ContainerData } from "$lib/types/ContainerData";
  import { onMount } from "svelte";

  let menuVisible = $state(false);
  let menuX = $state(0);
  let menuY = $state(0);
  let menuData = $state({} as ContainerData);

  function openMenu(event: MouseEvent, data: ContainerData) {
    event.stopPropagation();

    menuVisible = true;
    menuData = data;

    const rect = (event.target as HTMLElement).getBoundingClientRect();
    menuX = event.clientX;
    menuY = rect.bottom + window.scrollY;
  }

  function closeMenu() {
    menuVisible = false;
  }

  function openContainer(appId: string) {
    window.open(`//${appId}.${location.host}`, "_blank");
    closeMenu();
  }

  // Close menu on outside click
  onMount(() => {
    const handler = (e: MouseEvent) => {
      if (menuVisible) closeMenu();
    };
    window.addEventListener("click", handler);
    return () => window.removeEventListener("click", handler);
  });
</script>

<div class="apps">
  {#each page.data.containers as container}
    <div class="app" id={container.id}>
      <div class="content">
        <div class="top">
          <button class="menu-opener" onclick={(e) => openMenu(e, container)}>
            +
          </button>
        </div>
        <a
          class="info"
          href={`//${container.id}.${page.url.host}`}
          target="_blank"
        >
          <img
            class="icon"
            src={container.icon ? container.icon : "https://placehold.co/60"}
            alt="An icon"
          />
          <span class="name">{container.name}</span>
        </a>
      </div>
    </div>
  {/each}
</div>

{#if menuVisible}
  <div class="menu" style="left: {menuX}px; top: {menuY}px;">
    <button class="item" onclick={() => openContainer(menuData.id)}>
      Open
    </button>
  </div>
{/if}

<style>
  .apps {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(131px, 1fr));
    gap: 20px;
    margin: 20px;
    width: 100%;
  }

  .app {
    background: white;
    border-radius: 10px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    padding: 20px;
    user-select: none;
    height: 80px;
    width: 92px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  .app > .content {
    position: relative;
    bottom: 5px;
  }

  .app > .content > .top > .menu-opener {
    margin-left: 90px;
    border-radius: 50px;
    border-color: #419fff;
    font-size: 14px;
    width: 25px;
    height: 25px;
  }

  .app > .content > .info {
    text-decoration: none;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .apps > .app > .content > .info > .icon {
    width: 60px;
    height: 60px;
  }

  .apps > .app > .content > .info > .name {
    font-size: 1.2em;
    display: block;
    color: #333;
    overflow: hidden;
    text-overflow: ellipsis;
    text-wrap: nowrap;
    height: 20px;
    width: 92px;
  }

  .apps > .app > .content > .info:hover > .name {
    color: green;
  }

  .menu {
    position: absolute;
    background: white;
    border: 1px solid #ccc;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    min-width: 150px;
    display: flex;
    flex-direction: column;
    gap: 1px;
    border-radius: 8px;
    user-select: none;
    z-index: 1000;
  }

  .menu > button.item {
    cursor: pointer;
    display: flex;
    background: none;
    border: none;
    width: 100%;
    padding: 10px 16px;
    text-align: left;
    font-size: 1em;
  }

  .menu > button.item:hover {
    background-color: lightgrey;
  }
</style>
