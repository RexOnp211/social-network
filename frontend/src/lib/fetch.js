export default async function FetchFromBackend(endPoint, options) {
  const url = process.env.API_URL || "http://localhost:8080";
  const res = await fetch(`${url}${endPoint}`, options);
  return res.json();
}
