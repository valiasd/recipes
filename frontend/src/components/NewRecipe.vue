<template>
  <h1>Add New Recipe</h1>
  <div class="container">
    <div class="row">
      <div class="col-10"></div>
      <div class="col-2">
        <router-link class="btn btn-primary" to="/recipes">Back</router-link>
      </div>
    </div>
    <div class="row">
      <div class="col-1"></div>
      <div class="col-5">
        <form @submit.prevent="addRecipe">
          <div class="row">
            <label for="floatingInput">Title</label>
            <input type="text" class="form-control" id="floatingTitle" placeholder="Title" v-model="title">
          </div>
          <div class="row">
            <label for="floatingInput">Description</label>
            <textarea class="form-control" id="floatingDescription" placeholder="Description" v-model="description"></textarea>
          </div>
          <div class="row" style="margin-top: 20px;">
            <button class="btn btn-primary w-100 py-2" type="submit">Add Recipe</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
<script setup>
import {ref} from 'vue';
import {useRouter} from 'vue-router';
import axios from "axios";

const title = ref('');
const description = ref('');
const router = useRouter()

const addRecipe = async () => {
  const result = await axios.post('/api/recipes', {
    title: title.value,
    description: description.value,
    author_id: window.__recipesStore.user.ID,
  }, {
    headers: {
      'Content-Type': 'application/json'
    }
  })

  if (result.status === 200) {
    await router.push(`/recipes`)
  } else {
    alert('Failed to add recipe')
  }

};
</script>