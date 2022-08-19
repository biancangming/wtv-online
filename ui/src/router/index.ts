import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import EditView from "../views/Edit.vue";
import Layout from "@/components/Layout/Index.vue";
const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      component: Layout,
      redirect: {
        name: "home",
      },
      children: [
        {
          path: "/home",
          name: "home",
          component: HomeView,
        },
        {
          path: "/edit",
          name: "edit",
          component: EditView,
        },
      ],
    },
  ],
});

export default router;
