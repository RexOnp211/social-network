"use client";
import FetchFromBackend from "@/lib/fetch";
import { useRouter } from "next/navigation";

export default function Login() {
  const router = useRouter();
  async function handleSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const email = formData.get("email");
    const password = formData.get("password");
    console.log(email, password);

    try {
      const res = await FetchFromBackend("/login", {
        method: "POST",
        body: formData,
        credentials: "include",
      });
      if (res.status === 200) {
        console.log("Login successful");

        // clear login info from local storage
        localStorage.setItem("userID", res.userId);
        localStorage.setItem("user", res.username);

        router.push("/");
      } else {
        throw new Error(await res.text());
      }
    } catch (error) {
      console.error("error logging in", error);
      alert(error)
    }
  }

  return (
    <div className="flex w-screen h-screen justify-center items-center">
      <div className="p-6 bg-primary rounded-lg w-96 h-auto">
        <h2 className="text-center text-2xl m-1 mb-3">Log in</h2>
        <form
          onSubmit={handleSubmit}
          className="flex flex-col justify-center items-center"
        >
          <div className="mb-4 w-full">
            <label htmlFor="email" className="block ">
              Email/Username:
            </label>
            <input
              type="text"
              id="email"
              name="email"
              required
              className="w-full p-2 rounded-lg"
            />
          </div>
          <div className="mb-4 w-full">
            <label htmlFor="password" className="block">
              Password:
            </label>
            <input
              type="password"
              id="password"
              name="password"
              required
              className="w-full p-2 rounded-lg"
            />
          </div>
          <button
            type="submit"
            className="bg-accent w-full text-white rounded-lg p-2 transition-colors hover:bg-accentDark"
          >
            Log in
          </button>
        </form>
        <p className="mt-4 flex-col text-center">
          Dont have an account?{" "}
          <a
            href="/register"
            className="underline text-accent transition-colors hover:text-accentDark"
          >
            Register here
          </a>
        </p>
      </div>
    </div>
  );
}
