import { lazy, Show, Suspense } from "solid-js";
import { HashRouter, Navigate } from "@solidjs/router";
import "./index.scss";
import Layout from "./views/components/Layout";

const routes = [
  {
    path: "/login",
    component: lazy(() => import("./views/Login")),
  },
  {
    path: "/home",
    component: lazy(() => import("./views/Home")),
  },
  {
    path: "/",
    component: () => <Navigate href="/home" />,
  },
];

export default function App() {
  return (
    <HashRouter
      root={(props) => (
        <Show
          when={
            !props.location.pathname.startsWith("/viewer") &&
            !props.location.pathname.startsWith("/login")
          }
          fallback={props.children}
        >
          <Layout>
            <Suspense>{props.children}</Suspense>
          </Layout>
        </Show>
      )}
    >
      {routes}
    </HashRouter>
  );
}
