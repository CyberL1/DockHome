<script lang="ts">
  import { invalidate } from "$app/navigation";
  import { page } from "$app/state";
  import Menu from "$lib/components/Menu";
  import type { ContainerData } from "$lib/types/ContainerData";
  import { onMount } from "svelte";

  // Invalidate data every 3 seconds
  onMount(() => {
    const interval = setInterval(() => {
      invalidate("/api/data");
    }, 3000);

    return () => clearInterval(interval);
  });

  let menu: Menu;
  let menuData: ContainerData;

  function openMenu(event: MouseEvent, data: ContainerData) {
    menuData = data;
    menu.open(event, data);
  }

  function openContainer(appId: string) {
    window.open(`//${appId}.${page.data.domain}`, "_blank");
    menu.close();
  }
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
          href={`//${container.alias || container.name}.${page.data.domain}`}
          target="_blank"
        >
          <img
            src={container.icon ? container.icon : "https://placehold.co/60"}
            alt="An icon"
          />
          <span>{container.name}</span>
        </a>
      </div>
    </div>
  {/each}
</div>


<Menu
  bind:this={menu}
  data={menuData}
  items={[
    {
      title: "Open",
      onclick: () => openContainer(menuData.alias || menuData.name),
    },
  ]}
/>

<style>
  .apps {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(131px, 1fr));
    gap: 20px;
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

  .app > .content > a {
    text-decoration: none;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .apps > .app > .content > a > img {
    width: 60px;
    height: 60px;
  }

  .apps > .app > .content > a > span {
    font-size: 1.2em;
    display: block;
    color: #333;
    overflow: hidden;
    text-overflow: ellipsis;
    text-wrap: nowrap;
    width: 95px;
  }

  .apps > .app > .content > a:hover > span {
    color: green;
  }
</style>
