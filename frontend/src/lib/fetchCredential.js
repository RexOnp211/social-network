// fetch login user name
import FetchFromBackend from "@/lib/fetch";

export default async function fetchCredential() {
  const credentialResponse = await FetchFromBackend(`/credential`, {
    method: "GET",
    credentials: "include",
  });
  if (!credentialResponse.ok) {
    console.log(`Failed to fetch credential`);
  }
  const credential = await credentialResponse.json();
  console.log("Credential: ", credential);

  return credential;
}
