<template>
<h2>Catalog example</h2>
<div id="catalog-container">
    <div v-for="card in cards" :key="card.id">
    <a v-bind:href="'http://localhost:8081/#/product/'+card.id">
      <div class="card">
        <img class="card-picture" :src="card.picturepath"/>
        <h3>{{ card.price }} ₽</h3>
        <h4>{{ card.title }}</h4>
      </div>
      </a>
    </div>
</div>
</template>

<script>
import axios from "axios";
import underscore from 'underscore';
export default {
  name: 'CatalogTemplate',
  data() {
    return {
      cards: [],
      initCardsQuantity: 30,
      nextCardsQuantity: 15,
      loadedCardsCounter: 0,
      cardsLimit: 60,
      cardsBefore: 0
    };
  },

  methods: {
    getCards(startRow, quantity) {
      var bodyFormData = new FormData();
      bodyFormData.append('quantity', quantity);
      bodyFormData.append('startrow', startRow);
      axios({
        method: 'post',
        url: 'http://localhost:1111/api/cards',
        data: bodyFormData
      })
       .then(response => {
         this.cards.push(...response.data.message);
         this.loadedCardsCounter += quantity;
      });
    },

    handleScroll() {
      if (document.documentElement.clientHeight + window.pageYOffset >= document.body.offsetHeight-300) {
        if (this.loadedCardsCounter >= this.cardsLimit) {
          window.removeEventListener('scroll', this.handleScroll);
          return;
        }
        this.getCards(this.cardsBefore+this.loadedCardsCounter-1, this.nextCardsQuantity);
      }
    },

    debouncedScroll() {}
  },

  mounted () {
    window.addEventListener('scroll', this.debouncedScroll);
  },

  beforeUnmount() {
    window.removeEventListener('scroll', this.debouncedScroll);
  },

  beforeMount() {
    this.debouncedScroll = underscore.debounce(this.handleScroll, 50);

    this.cardsBefore = this.$route.params.page*this.cardsLimit;
    if (this.cardsBefore>0) {
      this.getCards(this.cardsBefore-1, this.initCardsQuantity);
    } else {
      this.getCards(this.cardsBefore, this.initCardsQuantity);
    }
  }
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#catalog-container {
  display: flex;
  justify-content:space-evenly;
  flex-wrap: wrap;
  gap: 20px 20px;
}

.card {
  width: 200px;
  margin: 1em auto;
}

.card-picture {
  width: 200px;
}

h4 {
  color: rgb(150, 150, 150);
}

a:link {
  text-decoration: none;
  color:#2c3e50;
}

a:visited {
  text-decoration: none;
  color:#2c3e50;
}

a:hover {
  width:300px;
}
</style>
