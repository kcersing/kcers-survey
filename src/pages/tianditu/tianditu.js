(function () {
  if (!window.location.href.includes("xxx.cn"))
    return;
  const f = document.getElementsByTagName('script')[0];
  const j = document.createElement('script');
  j.async = true;
  j.id = 'tianditu';
  j.type = 'text/javascript';
  j.src = "https://api.tianditu.gov.cn/api?v=4.0&tk=你的应用Key";
  f.parentNode.insertBefore(j, f);
})();
