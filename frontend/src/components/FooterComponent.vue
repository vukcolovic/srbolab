<template>
  <footer class="d-flex align-items-center justify-content-center bg-primary">
    Copyright &copy; {{ new Date().getFullYear()}}
  </footer>
</template>

<style scoped>
  footer {
    position:absolute;
    bottom:0;
    width:100%;
    color: white;
    min-height: 3em;
    font-size: 0.8em;
  }
</style>