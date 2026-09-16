export interface HeatPoint {
  lng: string;
  lat: string;
}

export function parsePoints(
  data: HeatPoint[] | null | undefined,
): Array<{ lng: number; lat: number }> {
  if (!data) return [];
  return data
    .map((d) => ({ lng: parseFloat(d.lng), lat: parseFloat(d.lat) }))
    .filter((p) => Number.isFinite(p.lng) && Number.isFinite(p.lat));
}

export function computeHeatMax(data: Array<{ count: number }> | null | undefined): number {
  if (!data || data.length === 0) return 1;
  const max = data.reduce((m, d) => Math.max(m, d.count || 0), 0);
  return max > 0 ? max : 1;
}

export function toHeatmapSeriesData(
  data: Array<{ lng: string; lat: string; count: number }> | null | undefined,
): Array<[number, number, number]> {
  if (!data) return [];
  return data
    .map((d) => ({ lng: parseFloat(d.lng), lat: parseFloat(d.lat), count: d.count || 0 }))
    .filter((p) => Number.isFinite(p.lng) && Number.isFinite(p.lat))
    .map((p) => [p.lng, p.lat, p.count]);
}