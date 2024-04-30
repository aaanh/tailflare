import Readme from "@/app/README.mdx";

export default function Home() {
  return (
    <main className="p-4">
      <h1 className="text-primary">Tailflare</h1>
      <article className="prose dark:prose-invert prose-lg">
        <Readme></Readme>
      </article>
    </main>
  );
}
