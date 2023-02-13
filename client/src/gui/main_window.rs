use crate::ClientState;
use druid::{AppDelegate, Command, DelegateCtx, Env, Handled, Target, Widget, WindowHandle, WindowId};
use druid::widget::Flex;
use crate::gui::menu_sidebar;
use spdlog::prelude::*;

pub struct Delegate
{
    pub window_ids: Vec<WindowId>,
}

impl AppDelegate<ClientState> for Delegate {
    fn command(&mut self, _ctx: &mut DelegateCtx, _target: Target, cmd: &Command, _data: &mut ClientState, _env: &Env) -> Handled {
        if let Some(_) = cmd.get(druid::commands::QUIT_APP)
        {
            println!("exiting client");
            return Handled::Yes;
        }

        Handled::No
    }

    fn window_added(&mut self, id: WindowId, _handle: WindowHandle, _data: &mut ClientState, _env: &Env, _ctx: &mut DelegateCtx) {
        println!("Window added, id: {:?}", id);
        self.window_ids.push(id);
    }

    fn window_removed(&mut self, id: WindowId, _data: &mut ClientState, _env: &Env, _ctx: &mut DelegateCtx) {
        println!("Window removed, id: {:?}", id);
        if let Some(pos) = self.window_ids.iter().position(|x| *x == id) {
            self.window_ids.remove(pos);
        }
    }
}

pub fn build() -> impl Widget<ClientState> {
    debug!("building main window");
    Flex::row().with_flex_child(menu_sidebar::make(), 1.0)
}