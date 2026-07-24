import { Outlet } from "react-router-dom";
import { useDocumentTitle } from "../lib/useDocumentTitle";

export default function WritingPage() {
  useDocumentTitle("Writing - My Blog");

  return (
    <div>
      <section className="mb-6">
        <h1 className="font-serif text-4xl md:text-5xl font-bold text-charcoal mb-4 tracking-tight">Writing</h1>
      </section>
      <Outlet />
    </div>
  );
}
