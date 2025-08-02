<header>
  <nav class="navbar">
    <div class="left">
      <a href="/" class="logo">DockHome</a>
      <a href="https://github.com/CyberL1/DockHome" target="_blank">
        Github repository
      </a>
    </div>
  </nav>
</header>

<main>
  <slot />
</main>

<style>
  :global(body) {
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
      "Helvetica Neue", Arial, sans-serif;
    background: #f0f0f0;
  }

  main {
    margin: 20px;
  }

  header {
    background-color: #24292e;
    padding: 10px 20px;
  }

  .navbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin: 0 auto;
  }

  .navbar a {
    color: #ffffff;
    text-decoration: none;
    margin-right: 20px;
    font-size: 14px;
    line-height: 1.5;
  }

  .navbar > .left > a.logo {
    font-weight: bold;
    font-size: 16px;
  }

  .navbar a:hover {
    text-decoration: underline;
  }
</style>
