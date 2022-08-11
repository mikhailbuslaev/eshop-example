<template>
<h1>header example</h1>
<nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
</nav>
</template>

<script>
export default {
    name: 'HeaderTemplate'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>