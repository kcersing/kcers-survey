import React from 'react';

class MapComponent extends React.Component {
  componentDidMount() {
    // 创建地图实例
    var map = new T.Map('mapContainer', {
      center: new T.LngLat(116.404, 39.915),
      zoom: 5
    });
    //初始化地图对象
    map = new T.Map("mapDiv");
    //设置显示地图的中心点和级别
    map.centerAndZoom(new T.LngLat(116.40969, 38.89945), zoom);
    var lo = new T.Geolocation();
    fn = function (e) {
      if (this.getStatus() == 0){
        map.centerAndZoom(e.lnglat, 15)
        alert("获取定位坐标："+e.lnglat.lat + "," + e.lnglat.lng)
        var marker = new T.Marker(e.lnglat);
        map.addOverLay(marker);

      }
      if(this.getStatus() == 1){
        map.centerAndZoom(e.lnglat, e.level)
        alert("获取定位坐标："+e.lnglat.lat + "," + e.lnglat.lng)
        var marker = new T.Marker(e.lnglat);
        map.addOverLay(marker);
      }
    }
    lo.getCurrentPosition(fn);



  }





  var map;
  var zoom = 12;
  function onLoad() {
    //初始化地图对象
    map = new T.Map("mapDiv");
    //设置显示地图的中心点和级别
    map.centerAndZoom(new T.LngLat(116.40969, 38.89945), zoom);
    var lo = new T.Geolocation();
    fn = function (e) {
      if (this.getStatus() == 0){
        map.centerAndZoom(e.lnglat, 15)
        alert("获取定位坐标："+e.lnglat.lat + "," + e.lnglat.lng)
        var marker = new T.Marker(e.lnglat);
        map.addOverLay(marker);

      }
      if(this.getStatus() == 1){
        map.centerAndZoom(e.lnglat, e.level)
        alert("获取定位坐标："+e.lnglat.lat + "," + e.lnglat.lng)
        var marker = new T.Marker(e.lnglat);
        map.addOverLay(marker);
      }
    }
    lo.getCurrentPosition(fn);
  }










  render() {
    return (
      <div id="mapDiv" style={{ width: '100%', height: '100%' }}></div>
    );
  }
}

export default MapComponent;
