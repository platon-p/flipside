export function Header() {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "row",
        height: 80,
        width: "100%",
        justifyContent: 'space-between',
        alignItems: 'center'
      }}
    >
      <div style={{width: 52, height: 52, backgroundColor: '#777'}}></div>
      <div>
        <a href="/login">войти</a>
      </div>
    </div>
  );
}
