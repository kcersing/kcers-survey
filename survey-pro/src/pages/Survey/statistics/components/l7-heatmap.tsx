import { Button } from 'antd';
import { HeatmapLayer, RasterLayer, Scene } from '@antv/l7';
import { Map } from '@antv/l7-maps';
import React, { useEffect, useRef } from 'react';
import { computeBounds, toHeatmapSeriesData } from './heatmap-utils';

const TIANDITU_TK = '516e46ec670dc4149ad67ed5020d99fd';
const DEFAULT_CENTER: [number, number] = [104.19, 35.86];
const DEFAULT_ZOOM = 4;
const TILE_EXTENT = [-180, -85.05112877980659, 180, 85.0511287798066];

const tileUrl = (type: string) =>
  `https://t0.tianditu.gov.cn/DataServer?T=${type}&X={x}&Y={y}&L={z}&tk=${TIANDITU_TK}`;

export const L7HeatMap = (props: { data: any[] }) => {
  const { data } = props;
  const containerRef = useRef<HTMLDivElement>(null);
  const sceneRef = useRef<any>(null);
  const boundsRef = useRef<any>(null);

  const fitToData = () => {
    const scene = sceneRef.current;
    if (!scene) return;
    const bounds = boundsRef.current;
    if (!bounds) {
      scene.setCenter(DEFAULT_CENTER);
      scene.setZoom(DEFAULT_ZOOM);
      return;
    }
    if (bounds.single) {
      scene.setCenter([bounds.minLng, bounds.minLat]);
      scene.setZoom(12);
      return;
    }
    scene.fitBounds(
      [
        [bounds.minLng, bounds.minLat],
        [bounds.maxLng, bounds.maxLat],
      ],
      { padding: 40 },
    );
  };
  const fitToDataRef = useRef(fitToData);
  fitToDataRef.current = fitToData;

  useEffect(() => {
    if (!containerRef.current) return;
    const scene = new Scene({
      id: containerRef.current,
      map: new Map({
        style: 'blank',
        center: DEFAULT_CENTER,
        zoom: DEFAULT_ZOOM,
        minZoom: 3,
        maxZoom: 18,
      }),
    });
    sceneRef.current = scene;

    scene.addLayer(
      new RasterLayer({ zIndex: 1 }).source(tileUrl('vec_w'), {
        parser: { type: 'rasterTile', tileSize: 256, minZoom: 3, maxZoom: 18, extent: TILE_EXTENT },
      }),
    );
    scene.addLayer(
      new RasterLayer({ zIndex: 2 }).source(tileUrl('cva_w'), {
        parser: { type: 'rasterTile', tileSize: 256, minZoom: 3, maxZoom: 18, extent: TILE_EXTENT },
      }),
    );

    const bounds = computeBounds(data);
    const points = toHeatmapSeriesData(data).map(([lng, lat, count]) => ({ lng, lat, count }));
    boundsRef.current = bounds ? { ...bounds, single: points.length <= 1 } : null;

    if (points.length > 0) {
      const heatLayer = new HeatmapLayer({ zIndex: 3 })
        .source(points, { parser: { type: 'json', x: 'lng', y: 'lat' } })
        .shape('heatmap')
        .size('count', [0, 1])
        .style({
          intensity: 2,
          radius: 18,
          rampColors: {
            colors: [
              '#313695',
              '#4575b4',
              '#74add1',
              '#fee090',
              '#fdae61',
              '#f46d43',
              '#a50026',
            ],
            positions: [0, 0.25, 0.4, 0.6, 0.75, 0.9, 1],
          },
        });
      scene.addLayer(heatLayer);
    }

    scene.on('loaded', () => {
      fitToDataRef.current();
    });

    return () => {
      sceneRef.current = null;
      scene.destroy();
    };
  }, [data]);

  return (
    <div style={{ position: 'relative', width: '100%', height: '100%' }}>
      <div ref={containerRef} style={{ width: '100%', height: '100%' }} />
      <Button
        size="small"
        onClick={fitToData}
        style={{ position: 'absolute', top: 8, right: 8, zIndex: 10 }}
      >
        回到全貌
      </Button>
    </div>
  );
};
