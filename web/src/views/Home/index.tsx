import { createSignal, onMount, Show } from "solid-js";
import { useTitle } from "../../signal/title";

export default function Home() {
  const [isLoading, setIsLoading] = createSignal(true);
  const [_, setTitle] = useTitle();

  onMount(() => {
    setTitle("首页");
    nextPage(0);
  });

  const [page, setPage] = createSignal(1);
  function nextPage(offset: number) {
    setIsLoading(true);
    const pg = page() + offset;
    setPage(pg);
    setIsLoading(false);
  }

  return (
    <>
      <Show when={!isLoading()}>
        <div class="flex gap-1">
          <button onClick={() => nextPage(-1)}>prev</button>
          <span>{page()}</span>
          <button onClick={() => nextPage(1)}>next</button>
        </div>
      </Show>
    </>
  );
}
