"use client";

import FetchFromBackend from "@/lib/fetch";
import { useState } from "react";
import { useRouter } from "next/navigation";

export default function Register() {
  const [ErrorMsg, setErrorMsg] = useState(null)
  const router = useRouter();
  const OnSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.target);
    const res = await FetchFromBackend("/register", {
      method: "POST",
      headers: {},
      body: formData,
    });
    if (res.ok) {
    window.location.href = "/login"
    } else {
      alert(await res.text())
    }
  };
  return (
    <div className="flex w-screen h-screen justify-center items-center">
      <div className="p-6 bg-primary rounded-lg w-96 h-auto">
        <h2 className="text-center text-2xl m-1 mb-3">Register</h2>
        <form onSubmit={OnSubmit} className="grid grid-cols-2 gap-4">
          <div className="mb-4 w-full col-span-2">
            <label htmlFor="nickname" className="block">
              Nickname
            </label>
            <input
              type="text"
              id="nickname"
              name="nickname"
              className="w-full p-2 rounded-lg"
            />
          </div>
          <div className="mb-4 w-full">
            <label htmlFor="email" className="block">
              Email <span className="text-red-500">*</span>
            </label>
            <input
              type="email"
              id="email"
              name="email"
              required
              className="w-full p-2 rounded-lg"
            />
          </div>
          <div className="mb-4 w-full">
            <label htmlFor="password" className="block">
              Password <span className="text-red-500">*</span>
            </label>
            <input
              type="password"
              id="password"
              name="password"
              required
              className="w-full p-2 rounded-lg"
            />
          </div>
          <div className="mb-4 w-full">
            <label htmlFor="firstname" className="block">
              First Name <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              id="firstname"
              name="firstname"
              required
              className="w-full p-2 rounded-lg"
            />
          </div>
          <div className="mb-4 w-full">
            <label htmlFor="lastname" className="block">
              Last Name <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              id="lastname"
              name="lastname"
              required
              className="w-full p-2 rounded-lg"
            />
          </div>

          <div className="mb-4 w-full">
            <label htmlFor="dob" className="block">
              Date of Birth <span className="text-red-500">*</span>
            </label>
            <input
              type="date"
              id="dob"
              name="dob"
              required
              className="w-full p-2 rounded-lg"
            />
          </div>
          <div className="mb-4 w-full">
            <label htmlFor="avatar" className="block">
              Avatar
            </label>
            <input
              type="file"
              id="avatar"
              name="avatar"
              className="w-full p-2"
            />
          </div>
          <div className="mb-4 w-full col-span-2">
            <label htmlFor="aboutme" className="block">
              About Me
            </label>
            <textarea
              id="aboutme"
              name="aboutme"
              className="w-full p-2 rounded-lg"
            />
          </div>
          <button
            type="submit"
            className="transiton-colors ease-in hover:bg-accentDark bg-accent w-full text-white rounded-lg p-2 col-span-2"
          >
            Register
          </button>
        </form>
        <p className="mt-4 flex-col text-center">
          Already have an account?
          <a
            href="/login"
            className="underline text-accent transition-colors hover:text-accentDark"
          >
            Log in here
          </a>
        </p>
      </div>
    </div>
  );
}
