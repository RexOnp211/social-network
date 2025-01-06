"use client"
import FetchFromBackend from "@/lib/fetch";
import Image from "next/image";
import { useEffect } from "react";
import { useState } from "react";

// width and height are the quality of the image and size is the size of the image in html
export default function ProfileImage({
  alt,
  width,
  height,
  size,
  avatar,
  className,
}) {
  const [url, setUrl] = useState("")
  useEffect(() => {
  FetchFromBackend(avatar, {
    credentials: "include"
  })
  .then(response => response.blob())
  .then(imageBlob => {
    // Create an image element and set its source to the blob
    setUrl(URL.createObjectURL(imageBlob));
    console.log("URL", url)
  })
  .catch(error => console.error('Error fetching image:', error));
}, [])

  return (
    <div className="overflow-hidden">
      <Image
        rel="preload"
        src={url || "https://churchthemes.com/wp-content/uploads/2016/07/google-maps-oops-something-went-wrong.png"}
        width={width}
        height={height}
        alt={alt}
        className={className}
        priority={true}
      />
    </div>
  );
}
