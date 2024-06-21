<template>
  <div>
    <main class="form-signin">
      <div class="container">
        <div class="row">
          <div class="col-4"></div>
          <div class="col-4">
            <p class="alert alert-info" v-if="newUsername">Welcome {{ newUsername }}, login with your new account</p>
            <form @submit.prevent="login">
              <h1 class="h3 mb-3 fw-normal">Please sign in</h1>
              <div class="row ">
                <label for="floatingInput">Username</label>
                <input type="text" class="form-control" id="floatingInput" placeholder="username" v-model="username">
              </div>
              <div class="row">
                <label for="floatingPassword">Password</label>
                <input type="password" class="form-control" id="floatingPassword" placeholder="Password" v-model="password">
              </div>
              <div class="row" style="margin-top: 20px;">
                <button class="btn btn-primary w-100 py-2" type="submit">Sign in</button>
              </div>
            </form>
            <p>Don't have an account? <router-link to="/register">Register</router-link></p>
            <div class="col-4"></div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
<script setup>
import {ref} from 'vue';
import {useRoute, useRouter} from "vue-router";
import axios from "axios";
const route = useRoute()
const router = useRouter()
const username = ref('');
const password = ref('');

const newUsername = ref(route.query.registered || null);

const login = async () => {
  try {
    const result = await axios.post('/api/users/login', {
      username: username.value,
      password: password.value
    }, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    console.log(result)

    if (result?.data.ID) {
      window.__recipesStore = {
        user: result.data
      }

      await router.push(`/recipes`)
   } else {
      alert('Failed to login')
    }
  } catch (e) {
    alert('Failed to login')
  }
};
</script>
