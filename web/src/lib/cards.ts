// Gradient pairs for bento cards: [from, to]
const cardGradients: [string, string][] = [
  ["from-card-orange", "to-rose-600"], // warm sunset
  ["from-card-dark", "to-navy-700"], // deep night
  ["from-card-blue", "to-indigo-500"], // ocean
  ["from-card-teal", "to-emerald-500"], // forest
  ["from-card-purple", "to-indigo-600"], // aurora
  ["from-amber-400", "to-card-orange"], // golden
];

export function cardGradientClass(index: number): string {
  const [from, to] = cardGradients[index % cardGradients.length];
  return `bg-gradient-to-br ${from} ${to}`;
}

export function isLightCard(index: number): boolean {
  return index % cardGradients.length === 5; // golden is the light one
}

export function cardTextClass(index: number): string {
  return isLightCard(index) ? "text-charcoal" : "text-white";
}

export function cardSubtextClass(index: number): string {
  return isLightCard(index) ? "text-charcoal/70" : "text-white/60";
}

export function cardBadgeClass(index: number): string {
  return isLightCard(index)
    ? "bg-black/10 text-charcoal backdrop-blur-sm"
    : "bg-white/15 text-white/90 backdrop-blur-sm";
}

export function capitalize(s: string): string {
  if (!s) return s;
  return s[0].toUpperCase() + s.slice(1);
}

export function formatDate(iso: string): string {
  const d = new Date(iso);
  return d.toLocaleDateString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
  });
}
