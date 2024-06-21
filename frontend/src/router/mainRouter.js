// router/mainRouter.js
import LoginPage from "../components/LoginPage.vue";

import {createMemoryHistory, createRouter} from 'vue-router'
import RegisterPage from "@/components/RegisterPage.vue";
import RecipesPage from "@/components/RecipesPage.vue";
import NewRecipe from "@/components/NewRecipe.vue";
import SingleRecipe from "@/components/SingleRecipe.vue";

const routes = [
    {
        path: '/',
        component: LoginPage,
        alias: '/login',
        beforeEnter: (to, from, next) => {
            if (!window.__recipesStore?.user) {
                next();
            } else {
                next('/recipes');
            }
        }
    },
    {
        path: '/recipes',
        component: RecipesPage,
        beforeEnter: (to, from, next) => {
            if (window.__recipesStore?.user) {
                next();
            } else {
                next('/login');
            }
        }

    },
    {
        path: '/recipes/new',
        component: NewRecipe,
        beforeEnter: (to, from, next) => {
            if (window.__recipesStore?.user) {
                next();
            } else {
                next('/login');
            }
        }

    },
    {
        path: '/recipes/:id',
        component: SingleRecipe,
        beforeEnter: (to, from, next) => {
            if (window.__recipesStore?.user) {
                next();
            } else {
                next('/login');
            }
        }

    },
    {
        path: '/register',
        component: RegisterPage,
        beforeEnter: (to, from, next) => {
            if (!window.__recipesStore?.user) {
                next();
            } else {
                next('/recipes');
            }
        }

    },
]

export const router = createRouter({
    history: createMemoryHistory(),
    routes,
})
