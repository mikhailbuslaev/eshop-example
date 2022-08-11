<template>
<h1>e-shop catalog example</h1>
<div id="catalog-body">
    <div v-for="card in cards" :key="card.id">
    <a v-bind:href="'http://localhost:8081/#/catalog/'+card.id">
      <div class="card">
        <img class="card-picture" :src="card.picturepath"/>
        <h3>{{ card.title }}</h3>
        <h4>{{ card.price }} ₽</h4>
      </div>
      </a>
    </div>
</div>
</template>

<script>
import axios from "axios";
export default {
  name: 'CatalogTemplate',
  data() {
    return {
      cards: [],
    };
  },

  methods: {
    getInitialCards() {
    var bodyFormData = new FormData();
    bodyFormData.append('quantity', 20);
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

    getNextCards(startRow = 0) {
      var bodyFormData = new FormData();
      bodyFormData.append('quantity', 10);
      bodyFormData.append('startrow', startRow);
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
  width: 180px;
  background: #ccc;
  margin: 1em auto;
}

.card-picture {
  width: 200px;
}
</style>
