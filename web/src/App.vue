<template>
  <div id="app">
    <main>
      <b-navbar toggleable="md" type="dark" variant="dark">
        <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
        <b-navbar-brand to="/">chi-vue-spa</b-navbar-brand>
        <b-collapse is-nav id="nav_collapse">
          <b-navbar-nav class="ml-auto">
            <b-nav-item to="/">Home</b-nav-item>
            <b-nav-item to="/data">Data</b-nav-item>
          </b-navbar-nav>
        </b-collapse>
      </b-navbar>
      <br />
      <b-container>
        <transition name="slidedown">
          <router-view />
        </transition>
      </b-container>
    </main>
    <b-footer class="text-muted">
      <div class="container">
        <p class="text-center">
          The source is available on
          <a href="https://github.com/zooraze/chi-vue-spa">GitHub</a> under the MIT license.
        </p>
      </div>
    </b-footer>
  </div>
</template>
