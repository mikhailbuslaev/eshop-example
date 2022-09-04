<template>
<div id="catalognavigation-container">
    <form @submit="reloadCatalog">
        <label for="category">
            <h3>Category</h3>
        </label>
        <input list="category" v-model="category">
        <datalist id="category">
            <option value="shoes"/>
        </datalist>
        <input type="submit" value="Search">
    </form>
</div>
</template>

<script>
export default {
  name: 'CatalogNavigation',
  data() {
    return {
      category: ''
    };
  },

  methods: {
    reloadCatalog() {
        this.$router.push({ path: '/catalog/'+this.category+'/0' });
    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>