<template>
  <div class="container">
    <h1>Recipes</h1>
    <div class="row">
      <div class="col-9"></div>
      <div class="col-3">
        <RouterLink class="btn btn-primary" to="/recipes/new">New Recipe</RouterLink>
      </div>
    </div>
      <div class="row" style="margin-top:20px;">
      <div class="col-1"></div>
      <div class="col-9">
        <table class="table table-bordered">
          <thead>
            <tr>
              <th scope="col">Title</th>
              <th scope="col">Description</th>
              <th scope="col" class="t-actions">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr class="trow" v-for="recipe in recipes" :key="recipe.ID">
              <td>{{ recipe.title }}</td>
              <td>{{ recipe.description }}</td>
              <td class="t-actions">
                <RouterLink class="btn btn-primary" :to="`/recipes/${recipe.ID}`">View</RouterLink>
                <button style="margin-left: 20px;" class="btn btn-danger" @click="deleteRecipe(recipe.ID)">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
<script setup>
import {onMounted, ref} from 'vue';
import axios from 'axios';

const recipes = ref([]);
onMounted(async () => {
  const recipesApi = await axios.get(`/api/recipes/author/${window.__recipesStore.user.ID}`);
  recipes.value = recipesApi.data;
})

const deleteRecipe = async (recipeId) => {
  const result = await axios.delete(`/api/recipes/${recipeId}`);
  if (result.status === 204) {
    recipes.value = recipes.value.filter(recipe => recipe.ID !== recipeId);
  } else {
    alert('Failed to delete recipe');
  }
}

</script>
<style>
.t-actions {
  text-align: right;
}
</style>