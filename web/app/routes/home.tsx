import type { Route } from "./+types/home";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "CV MAKER MOFO" },
    {
      name: "description",
      content: "Parce que chercher un job ça casse les couilles",
    },
  ];
}

export default function Home() {
  return <h1>COUCOU TOI, tu veux voir mon CV ?</h1>;
}
