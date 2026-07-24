export default function Footer() {
  const year = new Date().getFullYear();
  return (
    <footer className="mt-14 pt-6">
      <div className="flex flex-col sm:flex-row justify-between items-center gap-3 text-[11px] text-charcoal-light/30 tracking-wide">
        <p>&copy; {year} My Blog</p>
        <div className="flex items-center gap-4">
          <a href="#" className="hover:text-charcoal-light/60 transition-colors duration-200">
            GitHub
          </a>
          <span className="w-1 h-1 rounded-full bg-charcoal-light/15" />
          <a href="#" className="hover:text-charcoal-light/60 transition-colors duration-200">
            Twitter
          </a>
        </div>
      </div>
    </footer>
  );
}
