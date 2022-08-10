<template>
<h1>e-shop catalog example</h1>
<div id="catalog-body">
    <div class="card" v-for="card in cards" :key="card.id">
      <img class="card-picture" :src="card.picturepath"/>
      <h3>{{ card.title }}</h3>
    </div>
</div>
</template>

<script>
import axios from "axios";
export default {
  name: 'CatalogBody',
  el: '#catalog-body',
  data() {
    return {
      cards: [],
    };
  },

  methods: {
    getInitialCards() {
    var bodyFormData = new FormData();
    bodyFormData.append('quantity', 10);
    bodyFormData.append('startrow', 0);
    axios({
      method: 'post',
      url: 'http://localhost:1111/api/cards',
      data: bodyFormData
    })
    .then((response) => {
        this.cards = response.data.message;
      });
    },

    getNextCards() {
      var bodyFormData = new FormData();
      bodyFormData.append('quantity', 5);
      bodyFormData.append('startrow', 0);
      axios({
        method: 'post',
        url: 'http://localhost:1111/api/cards',
        data: bodyFormData
      })
       .then(response => {
         this.cards.push(...response.data.message);
      });
    },

    handleScroll() {
      if (window.innerHeight + window.pageYOffset >= document.body.offsetHeight) {
        this.getNextCards();
      }
    }
  },

  mounted () {
    window.addEventListener('scroll', this.handleScroll);
  },

  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
  },

  beforeMount() {
    this.getInitialCards();
  }
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#catalog-body {
    display: flex;
    flex-wrap: wrap;
    gap: 10px 10px;
}

.card {
  width: 200px;
  background: #ccc;
  border-radius: 1em;
  margin: 1em auto;
}

.card-picture {
  width: 200px;
}
</style>
