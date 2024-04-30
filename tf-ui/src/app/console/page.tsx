"use client";

import ConsoleForm from "./(components)/console-form";

export default function ConsolePage() {
  return (
    <main className="flex flex-col space-y-4">
      <h1 className="text-primary">Tailflare {">"} Console</h1>
      <ConsoleForm></ConsoleForm>
    </main>
  );
}
