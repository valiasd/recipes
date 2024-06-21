<template>
  <div class="container">
    <h1>{{ recipe.title }}</h1>
    <div class="row">
      <div class="col-10">
        <p>{{ recipe.description }}</p>
      </div>
      <div class="col-2">
        <router-link class="btn btn-primary" to="/recipes">Back</router-link>
      </div>
    </div>
  </div>
</template>
<script setup>
import {ref, onMounted} from 'vue';
import axios from 'axios';
import {useRoute} from 'vue-router';

const route = useRoute();
const recipe = ref({});
onMounted(async () => {
  const recipeApi = await axios.get(`/api/recipes/${route.params.id}`);
  recipe.value = recipeApi.data;
})
</script>