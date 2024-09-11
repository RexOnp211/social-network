export default async function Home() {
  const res = await fetch("http://localhost:8080/posts");
  const posts = await res.json();
  return (
    <>
      <h1>Home Page </h1>
      <p>{posts.message}</p>
    </>
  );
}
