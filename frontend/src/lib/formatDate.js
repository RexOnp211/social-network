// 1990-01-01T00:00:00Z -> 1990/01/01

export default function formatDate(isoDate) {
  const date = new Date(isoDate);

  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0"); // +1 since month starts from 0
  const day = String(date.getDate()).padStart(2, "0");

  return `${year}/${month}/${day}`;
}
