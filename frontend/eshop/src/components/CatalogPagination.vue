<template>
<div id="pagination-body" ref=paginationBody>
    <div v-for="paginationElement in paginationElements" :key="paginationElement.num">
      <a v-if="paginationElement.isLink" v-bind:href="'http://localhost:8081/#/catalog/'+paginationElement.num">
        <div class="pagination-element">
          <h3 v-if="paginationElement.isLink" >{{ paginationElement.num }}</h3>
          <h3 v-else>...</h3>
        </div>
        </a>
    </div>
</div>
</template>

<script>
import axios from "axios";
export default {
  name: 'CatalogPagination',
  data() {
    return {
        totalPagesCount: 0,
        cardsLimit: 25,
        currentPage: 0,
        paginationElements: []
    };
  },

  methods: {
    getPagination() {
      axios({
        method: 'post',
        url: 'http://localhost:1111/api/cards_count'
      })
       .then(response => {
         let totalCards = response.data.message;
         let lastCards = totalCards % this.cardsLimit;

         this.totalPagesCount = Math.floor(totalCards / this.cardsLimit);
         if (lastCards > 0) {
            this.totalPagesCount++;
         }

         if (this.totalPagesCount >= 7) {
            for (let i = 0; i < 7; i++) {
                this.paginationElements.push({num:0, isLink:true});
            }
            this.renderLongPagination();
            } else {
                for (let i = 0, length = this.totalPagesCount; i < length; i++) {
                    this.paginationElements.push({num:i, isLink:true});
                }
            }
        });
    },

    renderLongPagination() {
        if (this.currentPage < 5) {
            for (let i = 0; i < 5; i++) {
                this.paginationElements[i] = {num:i, isLink:true};
            }
        } else {
            this.paginationElements[0] = {num:0, isLink:true};
            this.paginationElements[1] = {num:0, isLink:false};
            this.paginationElements[2] = {num:this.currentPage-1, isLink:true};
            this.paginationElements[3] = {num:this.currentPage, isLink:true};
        }

        if (this.totalPagesCount-this.paginationElements[3] < 4){
            for (let i = 1; i < 4; i++) {
                this.paginationElements[i+3] = {num:this.paginationElements[3]+i, isLink:true};
            }
        } else {
            this.paginationElements[4] = {num:0, isLink:true};
            this.paginationElements[4] = {num:this.paginationElements[3]+1, isLink:false};
            this.paginationElements[5] = {num:0, isLink:false};
            this.paginationElements[6] = {num:this.totalPagesCount, isLink:true};
        }
    }
  },

  beforeMount() {
    this.currentPage = this.$route.params.page;
    this.getPagination();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#pagination-body {
    display: flex;
    flex-wrap: nowrap;
    justify-content: center;
}

.pagination-element {
  width: 40px;
  height: 40px;
  border: none;
  border-radius: 50%;
  background-color: #ccc;
  font-size: 18px;
  color: #2c3e50;
  margin: 20px;
}
</style>