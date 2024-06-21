<template>
  <div>
    <main class="form-signup">
      <div class="container">
        <div class="row">
          <div class="col-4"></div>
          <div class="col-4">
            <form @submit.prevent="register">
              <h1 class="h3 mb-3 fw-normal">Please sign in</h1>
              <div class="row">
                <label for="floatingInput">Username</label>
                <input type="text" class="form-control" id="floatingUsername" placeholder="username" v-model="username">
              </div>
              <div class="row">
                <label for="floatingInput">Email address</label>
                <input type="email" class="form-control" id="floatingInput" placeholder="name@example.com" v-model="email">
              </div>
              <div class="row">
                <label for="floatingPassword">Password</label>
                <input type="password" class="form-control" id="floatingPassword" placeholder="Password" v-model="password">
              </div>
              <div class="row" style="margin-top: 20px;">
                <button class="btn btn-primary w-100 py-2" type="submit">Register</button>
              </div>
            </form>
            <p>Already have an account? <router-link to="/">Login</router-link></p>
            <div class="col-4"></div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
<script setup>
import {ref} from 'vue';
import axios from "axios";
import { useRouter } from 'vue-router';

const username = ref('');
const email = ref('');
const password = ref('');
const router = useRouter()

const register = async () => {
  const result = await axios.post('/api/users/register', {
    username: username.value,
    email: email.value,
    password: password.value
  }, {
    headers: {
      'Content-Type': 'application/json'
    }
  })

  if (result.status === 200) {
    await router.push(`/?registered=${username.value}`)
  } else {
    alert('Failed to register')
  }

};
</script>
