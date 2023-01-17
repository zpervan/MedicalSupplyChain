fn main() {
    let parent_dir = std::env::current_dir().unwrap().parent().unwrap().to_owned();
    let proto_path = format!("{}{}", parent_dir.to_str().unwrap(), "/messages");
    let proto_file_path = format!("{}{}", proto_path, "/medical_supplies.proto");

    protoc_rust_grpc::Codegen::new()
        .out_dir("src")
        .include(proto_path)
        .input(proto_file_path)
        .rust_protobuf(true)
        .run()
        .expect("error compiling protocol buffer");
}