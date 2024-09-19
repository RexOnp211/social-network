import Image from "next/image";

export default function ProfileImage({ src, alt, width, height, size }) {
  return (
    <div className="overflow-hidden">
      <div style={{ width: size, height: size }}>
        <Image
          src={src}
          width={width}
          height={height}
          className="object-cover h-full w-full rounded-full"
          alt={alt}
          priority={true}
        />
      </div>
    </div>
  );
}
