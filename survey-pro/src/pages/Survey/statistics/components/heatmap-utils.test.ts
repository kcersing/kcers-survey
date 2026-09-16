import { computeHeatMax, parsePoints, toHeatmapSeriesData } from './heatmap-utils';

describe('parsePoints', () => {
  it('returns empty array for null/undefined/empty data', () => {
    expect(parsePoints(null as any)).toEqual([]);
    expect(parsePoints(undefined)).toEqual([]);
    expect(parsePoints([])).toEqual([]);
  });

  it('parses lng/lat strings to numbers', () => {
    const data = [
      { lng: '108.5', lat: '34.2' },
      { lng: '110.1', lat: '35.8' },
    ];
    expect(parsePoints(data)).toEqual([
      { lng: 108.5, lat: 34.2 },
      { lng: 110.1, lat: 35.8 },
    ]);
  });

  it('filters out non-numeric values', () => {
    const data = [
      { lng: 'abc', lat: '34.2' },
      { lng: '108.5', lat: '' },
      { lng: '109.0', lat: '33.9' },
    ];
    expect(parsePoints(data)).toEqual([{ lng: 109.0, lat: 33.9 }]);
  });
});

describe('computeHeatMax', () => {
  it('returns 1 for empty data', () => {
    expect(computeHeatMax([])).toBe(1);
    expect(computeHeatMax(null as any)).toBe(1);
  });

  it('returns the max count', () => {
    expect(computeHeatMax([{ count: 3 }, { count: 10 }, { count: 7 }])).toBe(10);
  });

  it('returns 1 when all counts are zero', () => {
    expect(computeHeatMax([{ count: 0 }])).toBe(1);
  });
});

describe('toHeatmapSeriesData', () => {
  it('returns empty array for null/undefined/empty data', () => {
    expect(toHeatmapSeriesData(null as any)).toEqual([]);
    expect(toHeatmapSeriesData(undefined)).toEqual([]);
    expect(toHeatmapSeriesData([])).toEqual([]);
  });

  it('converts lng/lat/count to [lng, lat, count] tuples', () => {
    const data = [
      { lng: '108.5', lat: '34.2', count: 2 },
      { lng: '110.1', lat: '35.8', count: 1 },
    ];
    expect(toHeatmapSeriesData(data)).toEqual([
      [108.5, 34.2, 2],
      [110.1, 35.8, 1],
    ]);
  });

  it('filters out invalid coordinates and defaults missing count to 0', () => {
    const data = [
      { lng: 'abc', lat: '34.2', count: 5 },
      { lng: '108.5', lat: '', count: 5 },
      { lng: '109.0', lat: '33.9', count: 0 },
    ];
    expect(toHeatmapSeriesData(data)).toEqual([[109.0, 33.9, 0]]);
  });
});