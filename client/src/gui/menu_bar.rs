use crate::ClientState;
use druid::{Data, Env, Menu, WindowId};

pub fn make<T: Data>(_window: Option<WindowId>, _data: &ClientState, _env: &Env) -> Menu<T> {
    let base = Menu::empty();
    base.entry(file()).entry(druid::platform_menus::win::file::exit())
}

fn file<T: Data>() -> Menu<T> {
    Menu::new("File").entry(druid::platform_menus::win::file::exit())
}