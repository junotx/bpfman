#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BytecodeImage {
    #[prost(string, tag = "1")]
    pub url: ::prost::alloc::string::String,
    #[prost(int32, tag = "2")]
    pub image_pull_policy: i32,
    #[prost(string, optional, tag = "3")]
    pub username: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, optional, tag = "4")]
    pub password: ::core::option::Option<::prost::alloc::string::String>,
}
/// BytecodeLocation is either:
/// - Parameters to pull an eBPF program stored in an OCI container image.
/// - Local file path for an image.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BytecodeLocation {
    #[prost(oneof = "bytecode_location::Location", tags = "2, 3")]
    pub location: ::core::option::Option<bytecode_location::Location>,
}
/// Nested message and enum types in `BytecodeLocation`.
pub mod bytecode_location {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Location {
        #[prost(message, tag = "2")]
        Image(super::BytecodeImage),
        #[prost(string, tag = "3")]
        File(::prost::alloc::string::String),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct KernelProgramInfo {
    /// Name as seen by the kernel. May be truncated version of user enter
    #[prost(uint32, tag = "1")]
    pub id: u32,
    /// name since limited to 16 byte by kernel
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    #[prost(uint32, tag = "3")]
    pub program_type: u32,
    #[prost(string, tag = "4")]
    pub loaded_at: ::prost::alloc::string::String,
    #[prost(string, tag = "5")]
    pub tag: ::prost::alloc::string::String,
    #[prost(bool, tag = "6")]
    pub gpl_compatible: bool,
    #[prost(uint32, repeated, tag = "7")]
    pub map_ids: ::prost::alloc::vec::Vec<u32>,
    #[prost(uint32, tag = "8")]
    pub btf_id: u32,
    #[prost(uint32, tag = "9")]
    pub bytes_xlated: u32,
    #[prost(bool, tag = "10")]
    pub jited: bool,
    #[prost(uint32, tag = "11")]
    pub bytes_jited: u32,
    #[prost(uint32, tag = "12")]
    pub bytes_memlock: u32,
    #[prost(uint32, tag = "13")]
    pub verified_insns: u32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ProgramInfo {
    /// Name entered by user
    #[prost(string, tag = "1")]
    pub name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub bytecode: ::core::option::Option<BytecodeLocation>,
    #[prost(message, optional, tag = "3")]
    pub attach: ::core::option::Option<AttachInfo>,
    #[prost(map = "string, bytes", tag = "4")]
    pub global_data: ::std::collections::HashMap<
        ::prost::alloc::string::String,
        ::prost::alloc::vec::Vec<u8>,
    >,
    #[prost(uint32, optional, tag = "5")]
    pub map_owner_id: ::core::option::Option<u32>,
    #[prost(string, tag = "6")]
    pub map_pin_path: ::prost::alloc::string::String,
    #[prost(string, repeated, tag = "7")]
    pub map_used_by: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(map = "string, string", tag = "8")]
    pub metadata: ::std::collections::HashMap<
        ::prost::alloc::string::String,
        ::prost::alloc::string::String,
    >,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct XdpAttachInfo {
    #[prost(int32, tag = "1")]
    pub priority: i32,
    #[prost(string, tag = "2")]
    pub iface: ::prost::alloc::string::String,
    #[prost(int32, tag = "3")]
    pub position: i32,
    #[prost(int32, repeated, tag = "4")]
    pub proceed_on: ::prost::alloc::vec::Vec<i32>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TcAttachInfo {
    #[prost(int32, tag = "1")]
    pub priority: i32,
    #[prost(string, tag = "2")]
    pub iface: ::prost::alloc::string::String,
    #[prost(int32, tag = "3")]
    pub position: i32,
    #[prost(string, tag = "4")]
    pub direction: ::prost::alloc::string::String,
    #[prost(int32, repeated, tag = "5")]
    pub proceed_on: ::prost::alloc::vec::Vec<i32>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TracepointAttachInfo {
    #[prost(string, tag = "1")]
    pub tracepoint: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct KprobeAttachInfo {
    #[prost(string, tag = "1")]
    pub fn_name: ::prost::alloc::string::String,
    #[prost(uint64, tag = "2")]
    pub offset: u64,
    #[prost(bool, tag = "3")]
    pub retprobe: bool,
    #[prost(string, optional, tag = "4")]
    pub namespace: ::core::option::Option<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct UprobeAttachInfo {
    #[prost(string, optional, tag = "1")]
    pub fn_name: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(uint64, tag = "2")]
    pub offset: u64,
    #[prost(string, tag = "3")]
    pub target: ::prost::alloc::string::String,
    #[prost(bool, tag = "4")]
    pub retprobe: bool,
    #[prost(int32, optional, tag = "5")]
    pub pid: ::core::option::Option<i32>,
    #[prost(string, optional, tag = "6")]
    pub namespace: ::core::option::Option<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AttachInfo {
    #[prost(oneof = "attach_info::Info", tags = "2, 3, 4, 5, 6")]
    pub info: ::core::option::Option<attach_info::Info>,
}
/// Nested message and enum types in `AttachInfo`.
pub mod attach_info {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Info {
        #[prost(message, tag = "2")]
        XdpAttachInfo(super::XdpAttachInfo),
        #[prost(message, tag = "3")]
        TcAttachInfo(super::TcAttachInfo),
        #[prost(message, tag = "4")]
        TracepointAttachInfo(super::TracepointAttachInfo),
        #[prost(message, tag = "5")]
        KprobeAttachInfo(super::KprobeAttachInfo),
        #[prost(message, tag = "6")]
        UprobeAttachInfo(super::UprobeAttachInfo),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct LoadRequest {
    #[prost(message, optional, tag = "1")]
    pub bytecode: ::core::option::Option<BytecodeLocation>,
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    #[prost(uint32, tag = "3")]
    pub program_type: u32,
    #[prost(message, optional, tag = "4")]
    pub attach: ::core::option::Option<AttachInfo>,
    #[prost(map = "string, string", tag = "5")]
    pub metadata: ::std::collections::HashMap<
        ::prost::alloc::string::String,
        ::prost::alloc::string::String,
    >,
    #[prost(map = "string, bytes", tag = "6")]
    pub global_data: ::std::collections::HashMap<
        ::prost::alloc::string::String,
        ::prost::alloc::vec::Vec<u8>,
    >,
    #[prost(string, optional, tag = "7")]
    pub uuid: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(uint32, optional, tag = "8")]
    pub map_owner_id: ::core::option::Option<u32>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct LoadResponse {
    #[prost(message, optional, tag = "1")]
    pub info: ::core::option::Option<ProgramInfo>,
    #[prost(message, optional, tag = "2")]
    pub kernel_info: ::core::option::Option<KernelProgramInfo>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct UnloadRequest {
    #[prost(uint32, tag = "1")]
    pub id: u32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct UnloadResponse {}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListRequest {
    #[prost(uint32, optional, tag = "1")]
    pub program_type: ::core::option::Option<u32>,
    #[prost(bool, optional, tag = "2")]
    pub bpfd_programs_only: ::core::option::Option<bool>,
    #[prost(map = "string, string", tag = "3")]
    pub match_metadata: ::std::collections::HashMap<
        ::prost::alloc::string::String,
        ::prost::alloc::string::String,
    >,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListResponse {
    #[prost(message, repeated, tag = "3")]
    pub results: ::prost::alloc::vec::Vec<list_response::ListResult>,
}
/// Nested message and enum types in `ListResponse`.
pub mod list_response {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct ListResult {
        #[prost(message, optional, tag = "1")]
        pub info: ::core::option::Option<super::ProgramInfo>,
        #[prost(message, optional, tag = "2")]
        pub kernel_info: ::core::option::Option<super::KernelProgramInfo>,
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PullBytecodeRequest {
    #[prost(message, optional, tag = "1")]
    pub image: ::core::option::Option<BytecodeImage>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PullBytecodeResponse {}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetRequest {
    #[prost(uint32, tag = "1")]
    pub id: u32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetResponse {
    #[prost(message, optional, tag = "1")]
    pub info: ::core::option::Option<ProgramInfo>,
    #[prost(message, optional, tag = "2")]
    pub kernel_info: ::core::option::Option<KernelProgramInfo>,
}
/// Generated client implementations.
pub mod bpfd_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct BpfdClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl<T> BpfdClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> BpfdClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            BpfdClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn load(
            &mut self,
            request: impl tonic::IntoRequest<super::LoadRequest>,
        ) -> std::result::Result<tonic::Response<super::LoadResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/bpfd.v1.Bpfd/Load");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("bpfd.v1.Bpfd", "Load"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn unload(
            &mut self,
            request: impl tonic::IntoRequest<super::UnloadRequest>,
        ) -> std::result::Result<tonic::Response<super::UnloadResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/bpfd.v1.Bpfd/Unload");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("bpfd.v1.Bpfd", "Unload"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn list(
            &mut self,
            request: impl tonic::IntoRequest<super::ListRequest>,
        ) -> std::result::Result<tonic::Response<super::ListResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/bpfd.v1.Bpfd/List");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("bpfd.v1.Bpfd", "List"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn pull_bytecode(
            &mut self,
            request: impl tonic::IntoRequest<super::PullBytecodeRequest>,
        ) -> std::result::Result<
            tonic::Response<super::PullBytecodeResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/bpfd.v1.Bpfd/PullBytecode",
            );
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("bpfd.v1.Bpfd", "PullBytecode"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get(
            &mut self,
            request: impl tonic::IntoRequest<super::GetRequest>,
        ) -> std::result::Result<tonic::Response<super::GetResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/bpfd.v1.Bpfd/Get");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("bpfd.v1.Bpfd", "Get"));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod bpfd_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with BpfdServer.
    #[async_trait]
    pub trait Bpfd: Send + Sync + 'static {
        async fn load(
            &self,
            request: tonic::Request<super::LoadRequest>,
        ) -> std::result::Result<tonic::Response<super::LoadResponse>, tonic::Status>;
        async fn unload(
            &self,
            request: tonic::Request<super::UnloadRequest>,
        ) -> std::result::Result<tonic::Response<super::UnloadResponse>, tonic::Status>;
        async fn list(
            &self,
            request: tonic::Request<super::ListRequest>,
        ) -> std::result::Result<tonic::Response<super::ListResponse>, tonic::Status>;
        async fn pull_bytecode(
            &self,
            request: tonic::Request<super::PullBytecodeRequest>,
        ) -> std::result::Result<
            tonic::Response<super::PullBytecodeResponse>,
            tonic::Status,
        >;
        async fn get(
            &self,
            request: tonic::Request<super::GetRequest>,
        ) -> std::result::Result<tonic::Response<super::GetResponse>, tonic::Status>;
    }
    #[derive(Debug)]
    pub struct BpfdServer<T: Bpfd> {
        inner: _Inner<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    struct _Inner<T>(Arc<T>);
    impl<T: Bpfd> BpfdServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            let inner = _Inner(inner);
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for BpfdServer<T>
    where
        T: Bpfd,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/bpfd.v1.Bpfd/Load" => {
                    #[allow(non_camel_case_types)]
                    struct LoadSvc<T: Bpfd>(pub Arc<T>);
                    impl<T: Bpfd> tonic::server::UnaryService<super::LoadRequest>
                    for LoadSvc<T> {
                        type Response = super::LoadResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::LoadRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Bpfd>::load(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = LoadSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/bpfd.v1.Bpfd/Unload" => {
                    #[allow(non_camel_case_types)]
                    struct UnloadSvc<T: Bpfd>(pub Arc<T>);
                    impl<T: Bpfd> tonic::server::UnaryService<super::UnloadRequest>
                    for UnloadSvc<T> {
                        type Response = super::UnloadResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::UnloadRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Bpfd>::unload(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = UnloadSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/bpfd.v1.Bpfd/List" => {
                    #[allow(non_camel_case_types)]
                    struct ListSvc<T: Bpfd>(pub Arc<T>);
                    impl<T: Bpfd> tonic::server::UnaryService<super::ListRequest>
                    for ListSvc<T> {
                        type Response = super::ListResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Bpfd>::list(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = ListSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/bpfd.v1.Bpfd/PullBytecode" => {
                    #[allow(non_camel_case_types)]
                    struct PullBytecodeSvc<T: Bpfd>(pub Arc<T>);
                    impl<T: Bpfd> tonic::server::UnaryService<super::PullBytecodeRequest>
                    for PullBytecodeSvc<T> {
                        type Response = super::PullBytecodeResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::PullBytecodeRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Bpfd>::pull_bytecode(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = PullBytecodeSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/bpfd.v1.Bpfd/Get" => {
                    #[allow(non_camel_case_types)]
                    struct GetSvc<T: Bpfd>(pub Arc<T>);
                    impl<T: Bpfd> tonic::server::UnaryService<super::GetRequest>
                    for GetSvc<T> {
                        type Response = super::GetResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Bpfd>::get(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = GetSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", "12")
                                .header("content-type", "application/grpc")
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: Bpfd> Clone for BpfdServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: Bpfd> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(Arc::clone(&self.0))
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: Bpfd> tonic::server::NamedService for BpfdServer<T> {
        const NAME: &'static str = "bpfd.v1.Bpfd";
    }
}
