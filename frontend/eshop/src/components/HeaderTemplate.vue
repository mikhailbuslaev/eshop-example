<template>
<div class="header">
  <h1>header example</h1>
  <nav>
      <router-link to="/catalog/0">Home</router-link> |
      <router-link to="/about">About</router-link>
  </nav>
</div>
</template>

<script>
export default {
    name: 'HeaderTemplate'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.header {
  height:200px;
  background-color: #ccc;
}

h1 {
  padding-top:30px;
  margin:0px;
}

nav {
  padding: 40px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>