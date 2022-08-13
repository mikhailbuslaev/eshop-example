<template>
<div id="pageup-button" >
    <button v-if="!isHidden" v-on:click="moveToTop">&#129045;</button>
</div>
</template>

<script>
export default {
  name: 'PageUpButton',
  data() {
    return {
      isHidden: true
    };
  },

  methods: {
    handleButtonVisibility() {
      if (window.pageYOffset >= document.documentElement.clientHeight) {
        this.isHidden = false;
      } else {
        this.isHidden = true;
      }
    },

    moveToTop() {
        window.scrollTo({ top: 0, behavior: 'smooth' });
        this.isHidden = true;
    }
  },

  mounted() {
    window.addEventListener('scroll', this.handleButtonVisibility);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#pageup-button {
  position: sticky;
  top: 100px;
}

button {
  position:absolute;
  right:5%;
  width: 40px;
  height: 40px;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  color:white;
  background-color: #ccc;
  font-size: 50px;
}
</style>