<script>

  import "../app.css";
  import '@skeletonlabs/skeleton/themes/theme-skeleton.css';
  import '@skeletonlabs/skeleton/styles/all.css';
  import { AppBar, GradientHeading } from '@skeletonlabs/skeleton';
</script>

<AppBar>
    <svelte:fragment slot="lead"><div class="text-3xl"><GradientHeading tag="h1" direction="bg-gradient-to-r" from="from-surface-500" to="to-surface-500">MuxKit</GradientHeading></div></svelte:fragment>
    <svelte:fragment slot="trail"></svelte:fragment>
</AppBar>
<main>
    <slot/>
</main>